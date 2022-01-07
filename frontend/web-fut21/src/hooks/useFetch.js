import { useEffect, useRef, useState } from "react";

export const useFetch = (url) => {
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
    fetch(url, {
      headers: { "Content-Type": "application/json", "x-api-key": "12345678" },
    })
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
  }, [url]);
  return state;
};
