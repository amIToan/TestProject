import axios from "axios";
import moment from "moment"
export function useAuth() {
  const token = localStorage.getItem("token")
  return token;
}
export function getUserInfo() {
  let userInfo = localStorage.getItem("userInfor") ? JSON.parse(localStorage.getItem("userInfor")) : null;
  return userInfo;
}
const BASE_URL = "http://localhost:3000/api"
axios.defaults.withCredentials = true;
export const costumedRequest = axios.create({
  baseURL: BASE_URL,
  withCredentials: true,
  //credentials: "include",
});
costumedRequest.interceptors.request.use((config) => {
  if (config.url !== "/user/login") {
    const authToken = useAuth();
    const dateTime = JSON.parse(localStorage
      .getItem("userInfor")
    )?.access_token_expires_at?.split(".")[0];
    const timeExpired = moment(dateTime).toDate().getTime();
    if (authToken && new Date().getTime() <= timeExpired) {
      config.headers.Authorization = `Bearer ${authToken}`;
    } else {
      localStorage.removeItem("token");
      localStorage.removeItem("userInfor");
    }
  }
  return config;
});
costumedRequest.interceptors.response.use(
  (response) => response,
  (error) => {
    const status = error.response ? error.response.status : null;

    if (status === 401) {
      //Handle unauthorized access
      localStorage.removeItem("token");
      window.location.href = "/login";
    } 
    return Promise.reject(error);
  }
);

