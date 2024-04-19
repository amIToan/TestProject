import { Button, DatePicker, Form, Input, Modal, Select } from "antd";
import React, { useState, forwardRef, useImperativeHandle } from "react";
import { PlusOutlined } from "@ant-design/icons";
import moment from "moment";
import { costumedRequest } from "../utils/helpers";
const AuthorModal = forwardRef(
  ({ resetDataTable, editedData, setEditData }, ref) => {
    const { Option } = Select;
    const [form] = Form.useForm();
    const [open, setOpen] = useState(false);
    const [confirmLoading, setConfirmLoading] = useState(false);
    const showModal = () => {
      setOpen(true);
    };
    const handleCancel = () => {
      setEditData(null);
      onReset();
      setConfirmLoading(false);
      setOpen(false);
    };
    const onReset = () => {
      form.resetFields();
    };
    const onFinish = async (values) => {
      try {
        setConfirmLoading(true);
        const newObject = {
          ...values,
          dob: moment(values.dob.$d).format("YYYY-MM-DD"),
        };
        if (editedData) {
          const { data } = await costumedRequest.put("/authors/update", {
            ...newObject,
            id: editedData?.id,
          });
          if (data) {
            resetDataTable((previousData) => {
              const oldData = [...previousData].map((item) => {
                if (item.id === data.id) {
                  return { ...data, key: data.id };
                }
                return item;
              });
              return oldData;
            });
          }
        } else {
          const { data } = await costumedRequest.post("/author", newObject);
          if (data) {
            resetDataTable((previousData) => {
              const oldData = [...previousData];
              oldData.push({ ...data, key: data.id });
              return oldData;
            });
          }
        }
        handleCancel();
      } catch (error) {
        setConfirmLoading(false);
      }
    };
    const onFinishFailed = (errorInfo) => {
      console.log("Failed:", errorInfo);
    };

    // Expose the function via ref
    useImperativeHandle(ref, () => ({
      showModal,
    }));

    return (
      <>
        {" "}
        <Button type="primary" icon={<PlusOutlined />} onClick={showModal}>
          Add new author
        </Button>
        <Modal
          title="Add author"
          open={open}
          confirmLoading={confirmLoading}
          onCancel={handleCancel}
          footer={null}
          width="600px"
        >
          <Form
            name="authors"
            form={form}
            labelCol={{
              span: 5,
            }}
            wrapperCol={{
              span: 16,
            }}
            style={{
              maxWidth: 600,
              width: 600,
            }}
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            autoComplete="off"
            scrollToFirstError
          >
            <Form.Item
              label="First name"
              name="first_name"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input!",
                },
              ]}
            >
              <Input defaultValue={editedData?.first_name} />
            </Form.Item>
            <Form.Item
              label="Last name"
              name="last_name"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input!",
                },
              ]}
            >
              <Input defaultValue={editedData?.last_name} />
            </Form.Item>
            <Form.Item label="Email" name="email">
              <Input defaultValue={editedData?.email} />
            </Form.Item>
            <Form.Item
              label="Birthday"
              name={"dob"}
              rules={[{ required: true, message: "Please input!" }]}
            >
              <DatePicker
                format={"YYYY-MM-DD"}
                defaultValue={editedData?.dob}
              />
            </Form.Item>
            <Form.Item
              label="Gender"
              name={"gender"}
              rules={[{ required: true, message: "Please input!" }]}
            >
              <Select
                placeholder="select your gender"
                defaultValue={editedData?.gender}
              >
                <Option value="male">male</Option>
                <Option value="female">female</Option>
                <Option value="other">other</Option>
              </Select>
            </Form.Item>
            <Form.Item
              wrapperCol={{
                offset: 8,
                span: 16,
              }}
            >
              <Button
                danger
                style={{ marginRight: 10 }}
                onClick={handleCancel}
                htmlType="reset"
              >
                Cancel
              </Button>
              <Button type="primary" htmlType="submit">
                Submit
              </Button>
            </Form.Item>
          </Form>
        </Modal>
      </>
    );
  }
);
export default AuthorModal;
