import { useEffect, useRef, useState } from "react";
import PropTypes from "prop-types";

export const useFetch = (url, { method, body }) => {
  const isMounted = useRef(true);

  const [state, setState] = useState({
    page: 0,
    totalPage: 0,
    items: 0,
    totalItems: 0,
    players: [],
  });

  useEffect(() => {
    return () => {
      //console.log("demontado");
      isMounted.current = false;
    };
  }, []);

  useEffect(() => {    
    let json = {
      headers: {
        "Content-Type": "application/json",
        "x-api-key": "12345678",
      },
    };
    if (method === "POST") {
      json = {
        ...json,
        method,
        body,
      };
    } else {
      json = {
        ...json,
        method,
      };
    }
    fetch(url, json)
      .then((resp) => resp.json())
      .then((data) => {
        if (isMounted.current) {
          //console.log(data);
          setState({
            page: data.Page,
            totalPage: data.TotalPages,
            items: data.Items,
            totalItems: data.TotalItems,
            players: data.Players,
          });
        }
      })
      .catch(() => {
        console.log("error peticion");
        setState({
          page: 0,
          totalPage: 0,
          items: 0,
          totalItems: 0,
          players: [],
        });
      });
  }, [url, method, body]);
  return state;
};

useFetch.propTypes = {
  ur: PropTypes.string.isRequired,
  method: PropTypes.string.isRequired,
  body: PropTypes.any.isRequired,
};
