import React from "react";
import {
  AppstoreOutlined,
  TeamOutlined,
  BookFilled,
  SearchOutlined,
} from "@ant-design/icons";
import { Avatar, Input, Layout, Menu } from "antd";
import { Route, Routes, useNavigate } from "react-router-dom";
import { Books } from "./screens/books";
import { Authors } from "./screens/authors";
import Genres from "./screens/genres";
import PrivateRouter from "./components/PrivateRouter";
const { Header, Content, Footer, Sider } = Layout;

const items = [
  {
    key: "books",
    icon: React.createElement(BookFilled),
    label: `Books`,
  },
  {
    key: "authors",
    icon: React.createElement(TeamOutlined),
    label: `Authors`,
  },
  {
    key: "genres",
    icon: React.createElement(AppstoreOutlined),
    label: `Genres`,
  },
];
function App() {
  const navigate = useNavigate();
  const handleNavigation = ({ key }) => {
    console.log(key);
    navigate("/" + key);
  };
  return (
    <Layout hasSider>
      <Sider className="!fixed !left-0 !top-0 !bottom-0 overflow-auto">
        <div className="demo-logo-vertical" />
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={["books"]}
          items={items}
          onClick={handleNavigation}
        />
      </Sider>
      <Layout className="ml-[200px]">
        <Header className="!bg-[#1677ff] sticky top-0 z-50 w-full p-3 !leading-none">
          <div className="flex justify-between items-center w-full">
            <div>
              <Input
                addonBefore={<SearchOutlined />}
                placeholder="large size"
                className="!h-auto"
              />
            </div>
            <Avatar size={40}></Avatar>
          </div>
        </Header>
        <Content style={{ margin: "16px 16px 0", overflow: "initial" }}>
          <Routes>
            <Route
              path="/"
              element={
                <PrivateRouter>
                  <Books />
                </PrivateRouter>
              }
            />
            <Route
              path="/books"
              element={
                <PrivateRouter>
                  <Books />
                </PrivateRouter>
              }
            />
            <Route
              path="/authors"
              element={
                <PrivateRouter>
                  <Authors></Authors>
                </PrivateRouter>
              }
            />
            <Route
              path="/genres"
              element={
                <PrivateRouter>
                  <Genres />
                </PrivateRouter>
              }
            />
          </Routes>
        </Content>
        <Footer style={{ textAlign: "center" }}>
          Ant Design ©{new Date().getFullYear()} Created by Ant UED
        </Footer>
      </Layout>
    </Layout>
  );
}

export default App;
