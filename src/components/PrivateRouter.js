import React from "react";
import { useAuth } from "../utils/helpers";
import { Navigate } from "react-router-dom";

const PrivateRouter = ({ children }) => {
  let auth = useAuth();
  return true ? <>{children}</> : <Navigate to="/login" />;
};

export default PrivateRouter;
