import { Button, Form, Input, Modal } from "antd";
import React, { useState, forwardRef, useImperativeHandle } from "react";
import { PlusOutlined } from "@ant-design/icons";
import { costumedRequest } from "../utils/helpers";
const GenreModal = forwardRef(
  ({ resetDataTable, editedData, setEditData }, ref) => {
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
        };
        if (editedData) {
          const { data } = await costumedRequest.put("/genres/update", {
            ...newObject,
            genre_id: editedData?.genre_id,
          });
          if (data) {
            resetDataTable((previousData) => {
              const oldData = [...previousData].map((item) => {
                if (item.genre_id === data.genre_id) {
                  return { ...data, key: data.genre_id };
                }
                return item;
              });
              return oldData;
            });
          }
        } else {
          const { data } = await costumedRequest.post("/genre", newObject);
          if (data) {
            resetDataTable((previousData) => {
              const oldData = [...previousData];
              oldData.push({ ...data, key: data.genre_id });
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
        <Button type="primary" icon={<PlusOutlined />} onClick={showModal}>
          Add new genres
        </Button>
        <Modal
          title="Add new genres"
          open={open}
          confirmLoading={confirmLoading}
          onCancel={handleCancel}
          footer={null}
          width="600px"
          destroyOnClose={true}
        >
          <Form
            name="genres"
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
              label="name"
              name="name"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input!",
                },
              ]}
            >
              <Input defaultValue={editedData?.name}/>
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
export default GenreModal;
