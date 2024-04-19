import { Button, Popconfirm, Space, Table, Tag } from 'antd';
import Column from 'antd/es/table/Column';
import ColumnGroup from 'antd/es/table/ColumnGroup';
import React, { useEffect, useRef, useState } from 'react'
import AuthorModal from '../components/AuthorModal';
import { costumedRequest } from '../utils/helpers';
import dayjs from 'dayjs';

export const Authors = () => {
   const childRef = useRef(null);
  const [tableData, setTableData] = useState([])
  const [editedData, setEditData] = useState(null)
  useEffect(() => {
    async function fetchData() {
       try {
         const { data: tableData } = await costumedRequest.get("/authors");
         if (tableData?.length > 0) {
           setTableData(tableData.map(item => ({...item, key: item.id})));
         }
       } catch (error) {
         console.log(error);
       }
    }
    fetchData();
  },[])
  const handleDeleteAuthor = async (id) => {
    try {
      const {data} = await costumedRequest.delete(`/authors/${id}`);
      if (data?.status == 200) {
        const filterData = tableData.filter((item) => item.id !== id )
        setTableData(filterData);
      }
    } catch (error) {
      console.error("err", error)
    }
  }
  const handleEditAuthor = (infor) => {
    if (infor && childRef.current) {
      const newInfo = {...infor, dob : dayjs(infor.dob, 'YYYY/MM/DD')}
      setEditData(newInfo);
       childRef.current.showModal();
    }
  }
  return (
    <>
      <div className="flex justify-end">
        <AuthorModal
          resetDataTable={setTableData}
          ref={childRef}
          editedData={editedData}
        />
      </div>
      <Table dataSource={tableData} bordered style={{ marginTop: 10 }}>
        <ColumnGroup title="Name">
          <Column title="First Name" dataIndex="first_name" key="first_name" />
          <Column title="Last Name" dataIndex="last_name" key="last_name" />
        </ColumnGroup>
        <Column title="Date Of Birth" dataIndex="dob" key="dob" />
        <Column title="Email" dataIndex="email" key="email" />
        <Column title="Gender" dataIndex="gender" key="gender" />
        <Column
          title="Action"
          key="action"
          render={(_, record) => (
            <Space size="middle">
              <Button
                onClick={() => {
                  handleEditAuthor(record);
                }}
              >
                Edit
              </Button>
              <Popconfirm
                title="Delete the task"
                description="Are you sure to delete this task?"
                okText="Yes"
                cancelText="No"
                onConfirm={() => {
                  handleDeleteAuthor(record.id);
                }}
              >
                <Button danger>Delete</Button>
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </>
  );
}
