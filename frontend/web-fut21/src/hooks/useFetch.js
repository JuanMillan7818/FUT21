import { useEffect, useRef, useState } from "react";

export const useFetch = (url, {method, body}) => {
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
      console.log("demontado");
      isMounted.current = false;
    };
  }, []);

  useEffect(() => {
    console.log("PEticion hecha a", url, method,body);
    let json = {
      headers: {
        "Content-Type": "application/json",
        "x-api-key": "12345678",
      }
    }
    if(method === 'POST') {
      json = {
        ...json,
        method,
        body,        
      }
    }else {
      json = {
        ...json,
        method,        
      }
    }    
    fetch(url, json)
      .then((resp) => resp.json())
      .then((data) => {
        if (isMounted.current) {
          console.log(data);
          setState({
            page: data.Page,
            totalPage: data.TotalPages,
            items: data.Items,
            totalItems: data.TotalItems,
            players: data.Players,
          });
        } else console.log("El set no se llamo");
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
