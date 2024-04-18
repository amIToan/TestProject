import React, { useState } from 'react'
import { Button, Checkbox, Form, Input } from "antd";
export const Login = () => {
  const [isLogin, setLoginForm] = useState(true)
  // const [loginState, setLoginStates] = useState({username: '', password:""})
  //  const [regisState, setRegisStates] = useState({
  //    username: "",
  //    password: "",
  //    email:"", 
  //    role:"admin"
  //  });
   
  const onFinish = (values) => {
    console.log("Success:", values);
  };
  const onFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };
  return (
    <div className="flex justify-center items-center w-full h-[100vh]">
      <Form
        name="login"
        labelCol={{
          span: 8,
        }}
        wrapperCol={{
          span: 16,
        }}
        style={{
          maxWidth: 600,
          width: 600,
        }}
        initialValues={{
          remember: true,
        }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
        scrollToFirstError
      >
        {isLogin ? (
          <>
            <Form.Item
              label="Username"
              name="username"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input your username!",
                },
              ]}
            >
              <Input
                // value={loginState.username}
                // onChange={(e) =>
                //   handleRegisLoginForm(e.target.value, "username")
                // }
              />
            </Form.Item>
            <Form.Item
              label="Password"
              name="password"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input your password!",
                },
              ]}
            >
              <Input.Password
                // value={loginState.password}
                // onChange={(e) =>
                //   handleRegisLoginForm(e.target.value, "password")
                // }
              />
            </Form.Item>
          </>
        ) : (
          <>
            <Form.Item
              label="Username"
              name="username"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input your username!",
                },
              ]}
            >
              <Input
                // value={regisState.username}
                // onChange={(e) =>
                //   handleRegisLoginForm(e.target.value, "username")
                // }
              />
            </Form.Item>
            <Form.Item
              label="Password"
              name="password"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input your password!",
                },
              ]}
            >
              <Input.Password
                // value={regisState.password}
                // onChange={(e) =>
                //   handleRegisLoginForm(e.target.value, "password")
                // }
              />
            </Form.Item>
            <Form.Item
              label="Email"
              name="email"
              validateTrigger="onBlur"
              rules={[
                {
                  type: "email",
                  message: "The input is not valid E-mail!",
                },
                {
                  required: true,
                  message: "Please input your E-mail!",
                },
              ]}
            >
              <Input
                // value={regisState.email}
                // onChange={(e) => handleRegisLoginForm(e.target.value, "email")}
              />
            </Form.Item>
          </>
        )}

        <Form.Item
          name="remember"
          valuePropName="checked"
          validateTrigger="onBlur"
          wrapperCol={{
            offset: 8,
            span: 16,
          }}
        >
          <div className="flex justify-between items-center">
            <Checkbox checked>Remember me</Checkbox>
            <a onClick={() => setLoginForm(!isLogin)}>
              {isLogin ? "Register" : "Login"}
            </a>
          </div>
        </Form.Item>

        <Form.Item
          wrapperCol={{
            offset: 8,
            span: 16,
          }}
        >
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}
