import React, { useState } from "react";
import { useFetch } from "../hooks/useFetch";
import { AppRouter } from "../routes/AppRouter";
import { ApiContext } from "./context/ApiContext";
import "animate.css";

export const FUT21 = () => {
  const [url, setUrl] = useState(`${process.env.REACT_APP_URL_PLAYER}page=1`);
  const [params, setParams] = useState({
    method: "GET",
    body: {},
  });
  const [searchExpand, setSearchExpand] = useState({});
  const { players, totalPage } = useFetch(url, params);

  return (
    <>
      <ApiContext.Provider
        value={{
          url,
          setUrl,
          players,
          totalPage,
          setParams,
          params,
          setSearchExpand,
          searchExpand,
        }}
      >
        <AppRouter />
      </ApiContext.Provider>
    </>
  );
};
