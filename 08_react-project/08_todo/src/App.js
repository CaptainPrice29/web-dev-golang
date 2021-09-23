import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { getData } from "./actions";
import TodoList from "./TodoList";
import "./App.css";

function App() {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(getData());
  }, [dispatch]);

  // const myState = useSelector((state) => state)
  // console.log("TESTING",myState);
  return (
    <div className="App">
      <TodoList></TodoList>
    </div>
  );
}

export default App;
