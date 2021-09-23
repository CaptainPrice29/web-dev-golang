import React, { useEffect, useRef, useState } from "react";
import { setData, getData } from "./actions";

import { useSelector, useDispatch } from "react-redux";
export default function TodoList() {
  const [sort, setsort] = useState("");
  const dispatch = useDispatch();
  const inputRefTitle = useRef(null);
  const inputRefDes = useRef(null);
  const myState = useSelector((state) => state.todoDataReducer);

  // let list = [1,2,3,4,5,6,7,'']
  let sourceElement = null;

  const [sortedList, setSortedList] = useState([]);
  useEffect(() => {
    myState.forEach((v) => {
      setSortedList(v);
    });
  }, [myState]);

  // sorting
  useEffect(() => {
    todoSort(sort);
  },[sort]);

  /* add a new entry at the end of the list.  */
  const newLine = () => {
    if (
      inputRefTitle.current.value !== "" &&
      inputRefDes.current.value !== ""
    ) {
      // console.log(sortedList);
      let newData = {
        createdAt: new Date().toString(),
        taskDescription: inputRefTitle.current.value,
        taskTitle: inputRefDes.current.value,
        status: true,
      };
      post(newData);
      setSortedList(sortedList.concat(newData));
      inputRefTitle.current.value = "";
      inputRefDes.current.value = "";
    }
  };

  /* change opacity for the dragged item. 
    remember the source item for the drop later */
  const handleDragStart = (event) => {
    event.target.style.opacity = 0.5;
    sourceElement = event.target;
    event.dataTransfer.effectAllowed = "move";
  };

  /* do not trigger default event of item while passing (e.g. a link) */
  const handleDragOver = (event) => {
    event.preventDefault();
    event.dataTransfer.dropEffect = "move";
  };

  /* add class .over while hovering other items */
  const handleDragEnter = (event) => {
    event.target.classList.add("over");
  };

  /* remove class .over when not hovering over an item anymore*/
  const handleDragLeave = (event) => {
    event.target.classList.remove("over");
  };

  const handleDrop = (event) => {
    /* prevent redirect in some browsers*/
    event.stopPropagation();

    /* only do something if the dropped on item is 
      different to the dragged item*/
    if (sourceElement !== event.target) {
      /* remove dragged item from list */
      const list = sortedList.filter(
        (item, i) => i.toString() !== sourceElement.id
      );

      /* this is the removed item */
      const removed = sortedList.filter(
        (item, i) => i.toString() === sourceElement.id
      )[0];

      /* insert removed item after this number. */
      let insertAt = Number(event.target.id);

      // console.log("list with item removed", list);
      // console.log("removed:  line", removed);
      // console.log("insertAt index", insertAt);

      let tempList = [];

      /* if dropped at last item, don't increase target id by +1. 
           max-index is arr.length */
      if (insertAt >= list.length) {
        tempList = list.slice(0).concat(removed);
        setSortedList(tempList);
        event.target.classList.remove("over");
      } else if (insertAt < list.length) {
        /* original list without removed item until the index it was removed at */
        tempList = list.slice(0, insertAt).concat(removed);

        // console.log("tempList", tempList);
        // console.log("insert the rest: ", list.slice(insertAt));

        /* add the remaining items to the list */
        const newList = tempList.concat(list.slice(insertAt));
        // console.log("newList", newList);

        /* set state to display on page */
        setSortedList(newList);
        event.target.classList.remove("over");
        // dispatch(setData(sortedList));
      }
    } else console.log("nothing happened");
    event.target.classList.remove("over");
  };

  const handleDragEnd = (event) => {
    event.target.style.opacity = 1;
  };

  // changes in current input field
  // const handleChange = (event) => {
  //   event.preventDefault();

  //   const list = sortedList.map((item, i) => {
  //     if (i !== Number(event.target.id)) {
  //       return item;
  //     } else return event.target.value;
  //   });
  //   setSortedList(list);
  // };

  // filter list only items with id unequal to current id
  const handleDelete = (event) => {
    event.preventDefault();
    const list = sortedList.filter((item, i) => i !== Number(event.target.id));
    // console.log(event.target.id);
    setSortedList(list);
  };
  // striking function
  function strikeThrough(text) {
    return text
      .split("")
      .map((char) => char + "\u0336")
      .join("");
  }
  //working with server dataTransfer
  async function post(newdata) {
    try {
      let res = await fetch(
        "http://6146ecde65467e00173849b9.mockapi.io/todoApi/task",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(newdata),
        }
      );
      dispatch(getData());
      console.log(res);
    } catch (error) {
      console.log(error);
    }
  }
  async function put(newdata, id) {
    try {
      let res = await fetch(
        "http://6146ecde65467e00173849b9.mockapi.io/todoApi/task/" + id,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ status: newdata }),
        }
      );
      console.log(res);
    } catch (error) {
      console.log(error);
    }
  }
  async function removedata(id) {
    try {
      let res = await fetch(
        "http://6146ecde65467e00173849b9.mockapi.io/todoApi/task/" + id,
        {
          method: "DELETE",
        }
      );
      console.log(res);
    } catch (error) {
      console.log(error);
    }
  }

  function todoSort(s) {
    switch (s) {
      case "all":
        {
          const t = [];
          myState.forEach((item) => {
            item.forEach((i) => {
              if (i.status === true) {
                t.unshift(i);
              } else {
                t.push(i);
              }
            });
          });
          setSortedList(t);
        }

        break;
      case "completed":
        {
          const t = [];
          myState.forEach((item) => {
            item.forEach((i) => {
              if (i.status === false) {
                t.push(i);
              }
            });
          });
          setSortedList(t);
        }

        break;
      case "uncompleted":
        {
          const t = [];
          myState.forEach((item) => {
            item.forEach((i) => {
              if (i.status === true) {
                t.push(i);
              }
            });
          });
          setSortedList(t);
        }

        break;

    }
  }

  // create list of items 
  const listItems = () => {
    return sortedList.map((item, i) => (
      <div key={i} className="todo-list">
        <input
          className="radio"
          type="radio"
          value={item.status}
          onClick={() => {
            item.status = false;
            put(false, item.id);
            strikeThrough(item.taskDescription);
            dispatch(setData(sortedList));
          }}
        />
        <li
          id={i}
          type="text"
          className="input-item"
          draggable="true"
          onDragStart={handleDragStart}
          onDragOver={handleDragOver}
          onDragEnter={handleDragEnter}
          onDragLeave={handleDragLeave}
          onDrop={handleDrop}
          onDragEnd={handleDragEnd}
          // onChange={handleChange}
          value={item.taskDescription}
        >
          {(() => {
            if (item.status === false) {
              return strikeThrough(
                item.taskDescription + " :- " + item.taskTitle
              );
            } else {
              return item.taskDescription + " :- " + item.taskTitle;
            }
          })()}
        </li>
        {/* /> */}
        <div
          id={i}
          className="delButton"
          onClick={(e) => {
            handleDelete(e);
            removedata(item.id);
            // dispatch(getData());
          }}
        >
          X
        </div>
      </div>
    ));
  };

  // console.log("sorted", sortedList);

  return (
    <div className="container">
      <h1 style={{ color: "white", textAlign: "center" }}>Todo-List</h1>
      <div id="setInput" className="form-group">
        <label
          style={{ color: "white", fontWeight: "bolder" }}
          htmlFor="formGroupExampleInput"
        >
          SET TITLE
        </label>
        <input
          type="text"
          className="form-control"
          id="formGroupExampleInput"
          ref={inputRefTitle}
        />
        <label
          style={{ color: "white", fontWeight: "bolder" }}
          htmlFor="formGroupExampleInput"
        >
          SET DESCRIPTION
        </label>
        <input
          type="text"
          className="form-control"
          id="formGroupExampleInput"
          ref={inputRefDes}
        />
        <button
          id="addbutton"
          className="addButton"
          onClick={() => {
            newLine();
          }}
        >
          ADD
        </button>
      </div>
      <select id="sort" onClick={(e) => setsort(e.target.value)}>
        <option id="t" value="all">all</option>
        <option value="completed">completed</option>
        <option value="uncompleted">uncompleted</option>
      </select>
      {listItems()}
    </div>
  );
}
