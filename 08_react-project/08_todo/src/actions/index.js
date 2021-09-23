export const getData = () => {
  return {
    type: "GET_DATA",
  };
};
export const setData = (data) => {
  return {
    type: "SET_DATA",
    data:data,
  };
};
