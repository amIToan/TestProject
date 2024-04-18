import { Button, Modal } from "antd";
import React, { useState } from "react";
import {
  PlusOutlined
} from "@ant-design/icons";
 //title, description, image_links, price, publication_date, author_id;
const BookModal = () => {
  const [open, setOpen] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [modalText, setModalText] = useState("Content of the modal");
  const handleOk = () => {
    setModalText("The modal will be closed after two seconds");
    setConfirmLoading(true);
    setTimeout(() => {
      setOpen(false);
      setConfirmLoading(false);
    }, 2000);
  };
  const showModal = () => {
    setOpen(true);
  };
  const handleCancel = () => {
    console.log("Clicked cancel button");
    setOpen(false);
  };
  return (
    <>
      {" "}
      <Button type="primary" icon={<PlusOutlined/>} onClick={showModal}>
        Add new book
      </Button>
      <Modal
        title="Add a new book"
        open={open}
        onOk={handleOk}
        confirmLoading={confirmLoading}
        onCancel={handleCancel}
      >
        <p>ahihi {modalText}</p>
      </Modal>
    </>
  );
};

export default BookModal;
