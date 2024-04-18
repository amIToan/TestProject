
import { Card } from 'antd';
import React from 'react'
import BookModal from '../components/BookModal';

export const Books = () => {
  const { Meta } = Card;
  return (
    <>
      <div className='flex justify-end'>
        <BookModal />
      </div>
      <div className="flex flex-wrap items-center w-full">
        <div className=" min-w-40 w-[25%] max-w-[25%] p-2">
          <Card
            hoverable
            className="shadow"
            cover={
              <img
                alt="example"
                src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
              />
            }
          >
            <Meta title="Europe Street beat" description="www.instagram.com" />
          </Card>
        </div>
        <div className=" min-w-40 w-[25%] max-w-[25%] p-2">
          <Card
            hoverable
            className="shadow"
            cover={
              <img
                alt="example"
                src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
              />
            }
          >
            <Meta title="Europe Street beat" description="www.instagram.com" />
          </Card>
        </div>
        <div className=" min-w-40 w-[25%] max-w-[25%] p-2">
          <Card
            hoverable
            className="shadow"
            cover={
              <img
                alt="example"
                src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
              />
            }
          >
            <Meta title="Europe Street beat" description="www.instagram.com" />
          </Card>
        </div>
      </div>
    </>
  );
}
