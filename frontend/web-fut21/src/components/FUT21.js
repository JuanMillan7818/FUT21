import React, { useState } from "react";
import { useFetch } from "../hooks/useFetch";
import { AppRouter } from "../routes/AppRouter";
import { ApiContext } from "./context/ApiContext";

export const FUT21 = () => {
  const [url, setUrl] = useState("http://localhost:8000/api/v1/players?page=1");
  const { players, totalPage } = useFetch(url);
  return (
    <>
      <ApiContext.Provider
        value={{
          url,
          setUrl,
          players,
          totalPage,
        }}
      >
        <AppRouter />
      </ApiContext.Provider>
    </>
  );
};
