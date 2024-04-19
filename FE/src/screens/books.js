import { Card } from "antd";
import React, { useContext, useEffect, useRef, useState } from "react";
import { DeleteFilled } from "@ant-design/icons";
import BookModal from "../components/BookModal";
import { costumedRequest } from "../utils/helpers";
import { AppContext } from "../App";
export const Books = () => {
  const context = useContext(AppContext);
  const { Meta } = Card;
  const childRef = useRef(null);
  const [books, setBooks] = useState([]);
  const [author, setAuthors] = useState();
  const [genres, setGenres] = useState();
  const [editedData, setEditData] = useState(null);
  async function fetchData() {
    let queryString = context.keywords ? `books/title?keyword=${context.keywords}` : "/books"; 
    try {
      const data = await Promise.all([
        costumedRequest.get(queryString),
        costumedRequest.get("/authors"),
        costumedRequest.get("/genres"),
      ]);
      if (data.length > 0) {
        setAuthors(data[1].data);
        setGenres(data[2].data);
      }
      const bookData = await Promise.all(
        [...data[0].data].map(async (element) => {
          const { data: book_genres } = await costumedRequest.get(
            `/books_genres/${element.book_id}`
          );
          const genderString = book_genres?.reduce(
            (accumulator, currentValue) =>
              accumulator + currentValue?.name + ", ",
            ""
          );
          return { ...element, book_genres: genderString };
        })
      );
      setBooks(bookData);
    } catch (error) {
      console.log(error);
    }
  }
  useEffect(() => {
    fetchData();
  }, [context.keywords]);
  const handleDeleteBook = async (id) => {
    if (!id) return;
    try {
      const { data } = await costumedRequest.delete(`/book/delete/${id}`);
       data?.status === 200 && fetchData();
    } catch (error) {
      console.error("err", error);
    }
  
  };
  return (
    <>
      <div className="flex justify-end">
        <BookModal
          resetBookData={setBooks}
          ref={childRef}
          editedData={editedData}
          authors={author}
          genres={genres}
          setEditData={setEditData}
          resetData={fetchData}
        />
      </div>
      <div className="flex flex-wrap items-stretch w-full">
        {books && books.length > 0 &&
          books.map((item, index) => (
            <div
              className="min-w-40 w-[25%] max-w-[25%] p-2 relative costumed_hover"
              key={index}
            >
              <div
                className="absolute text-white text-xl hidden text-red top-0 right-0 !z-50 close_button cursor-pointer bg-neutral-500 rounded-full w-[30px] h-[30px] justify-center items-center"
                onClick={() => handleDeleteBook(item.book_id)}
              >
                <DeleteFilled />
              </div>
              <Card
                hoverable
                className="shadow"
                cover={
                  <img
                    alt="example"
                    src={"http://localhost:3000/api" + item?.image_links[0]}
                  />
                }
              >
                <Meta title={item.title} />
                <div className="overflow-hidden line-clamp-1">
                  {item.description}
                </div>
                <div className="overflow-hidden line-clamp-1">
                  Price : {item.price}
                </div>
                <div>Author : {item.author_name}</div>
                <div> Publication : {item.publication_date}</div>
                <div> Genres : {item.book_genres}</div>
              </Card>
            </div>
          ))}
      </div>
    </>
  );
};
