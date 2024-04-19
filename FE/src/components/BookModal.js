import {
  Button,
  DatePicker,
  Form,
  Input,
  InputNumber,
  Modal,
  Select,
  Upload,
} from "antd";
import React, { forwardRef, useImperativeHandle, useState } from "react";
import { PlusOutlined } from "@ant-design/icons";
import TextArea from "antd/es/input/TextArea";
import moment from "moment";
import { costumedRequest } from "../utils/helpers";
//title, description, image_links, price, publication_date, author_id;
const BookModal = forwardRef(
  (
    { resetBookData, editedData, setEditData, genres, authors, resetData },
    ref
  ) => {
    const [form] = Form.useForm();
    const [open, setOpen] = useState(false);
    const [confirmLoading, setConfirmLoading] = useState(false);
    const [imgsStore, setImgStore] = useState([]);
    const showModal = () => {
      setOpen(true);
    };
    const handleCancel = () => {
      setImgStore([]);
      setEditData(null);
      onReset();
      setConfirmLoading(false);
      setOpen(false);
    };
    const onReset = () => {
      form.resetFields();
    };
    const onFinish = async (values) => {
      console.log("Success:", values);
      try {
        const newObject = {
          ...values,
          publication_date: moment(values.publication_date.$d).format(
            "YYYY-MM-DD"
          ),
        };
        const newFromData = new FormData();
        for (const [key, value] of Object.entries(newObject)) {
          if (key == "genre_id") {
            value.forEach((item) => {
              newFromData.append(key, item);
            });
          } else {
            newFromData.append(key, value);
          }
        }
        if (imgsStore.length > 0) {
          imgsStore[0].forEach((item) =>
            newFromData.append("image_links", item.originFileObj)
          );
        }
        const {
          data: { book: insertedBook },
        } = await costumedRequest.post("/book", newFromData, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        if (insertedBook) {
          await resetData();
          console.log("insertBook", insertedBook);
        }
        handleCancel();
      } catch (error) {
        console.error("error", error);
      }
    };
    const onFinishFailed = (errorInfo) => {
      console.log("Failed:", errorInfo);
    };
    const handleUploadImgs = ({ fileList }) => {
      if (fileList.length > 0) {
        setImgStore([fileList]);
        setConfirmLoading(false);
      }
    };
    // Expose the function via ref
    useImperativeHandle(ref, () => ({
      showModal,
    }));
    return (
      <>
        {" "}
        <Button type="primary" icon={<PlusOutlined />} onClick={showModal}>
          Add new book
        </Button>
        <Modal
          title="Add a new book"
          open={open}
          confirmLoading={confirmLoading}
          onCancel={handleCancel}
          width="600px"
          footer={null}
        >
          <Form
            name="books"
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
            initialValues={{
              remember: true,
            }}
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            autoComplete="off"
            scrollToFirstError
          >
            <Form.Item
              label="Title"
              name="title"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input title!",
                },
              ]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              label="Description"
              name="description"
              validateTrigger="onBlur"
              rules={[
                {
                  required: true,
                  message: "Please input your description !",
                },
              ]}
            >
              <TextArea rows={4} />
            </Form.Item>
            <Form.Item
              label="Upload"
              valuePropName="image_links"
              //getValueFromEvent={normFile}
              rules={[{ required: true, message: "Please input!" }]}
            >
              <Upload
                listType="picture-card"
                multiple
                onChange={handleUploadImgs}
                beforeUpload={() => {
                  return false;
                }}
                //fileList={imgsStore}
              >
                <button style={{ border: 0, background: "none" }} type="button">
                  <PlusOutlined />
                  <div style={{ marginTop: 8 }}>Upload</div>
                </button>
              </Upload>
            </Form.Item>
            <Form.Item label="Price" name="price">
              <InputNumber />
            </Form.Item>
            <Form.Item
              label="Publication Date"
              name={"publication_date"}
              rules={[{ required: true, message: "Please input!" }]}
            >
              <DatePicker />
            </Form.Item>
            <Form.Item
              label="Authors"
              name={"author_id"}
              rules={[{ required: true, message: "Please input!" }]}
            >
              <Select
                size="middle"
                placeholder="Please select"
                //onChange={handleChange}
                style={{ width: "100%" }}
                options={authors?.map((item) => {
                  return {
                    id: item.id,
                    label: item.first_name + " " + item.last_name,
                    value: item.id,
                  };
                })}
              />
            </Form.Item>
            <Form.Item
              label="Genres"
              name={"genre_id"}
              rules={[{ required: true, message: "Please input!" }]}
            >
              <Select
                mode="multiple"
                size="middle"
                placeholder="Please select"
                //onChange={handleChange}
                style={{ width: "100%" }}
                options={genres?.map((item) => {
                  return {
                    id: item.genre_id,
                    label: item.name,
                    value: item.genre_id,
                  };
                })}
              />
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

export default BookModal;
