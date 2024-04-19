import React, { useEffect, useRef, useState } from "react";
import { Button, Space, Table } from "antd";
import Column from "antd/es/table/Column";
import { costumedRequest } from "../utils/helpers";
import GenreModal from "../components/GenreModal";
export default function Genres(params) {
  const childRef = useRef(null);
  const [tableData, setTableData] = useState([]);
  const [editedData, setEditData] = useState(null);
  useEffect(() => {
    async function fetchData() {
      try {
        const { data: tableData } = await costumedRequest.get("/genres");
        if (tableData?.length > 0) {
          setTableData(
            tableData.map((item) => ({ ...item, key: item.genre_id }))
          );
        }
      } catch (error) {
        console.log(error);
      }
    }
    fetchData();
  }, []);

  const handleEditGenres = (infor) => {
    if (infor && childRef.current) {
      const newInfo = { ...infor};
      setEditData(newInfo);
      childRef.current.showModal();
    }
  };
  return (
    <>
      <div className="flex justify-end">
        <GenreModal
          resetDataTable={setTableData}
          ref={childRef}
          editedData={editedData}
          setEditData={setEditData}
        />
      </div>
      <Table dataSource={tableData} bordered style={{ marginTop: 10 }}>
        <Column title="STT" dataIndex="genre_id" key="genre_id" />
        <Column title="Name" dataIndex="name" key="name" />
        <Column
          title="Action"
          key="action"
          render={(_, record) => (
            <Space size="middle">
              <Button
                onClick={() => {
                  handleEditGenres(record);
                }}
              >
                Edit
              </Button>
            </Space>
          )}
        />
      </Table>
    </>
  );
}
