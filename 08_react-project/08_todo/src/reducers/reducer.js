const todoDataReducer = (state = [], action) => {
    switch (action.type) {
      case "SET_DATA":
        return [action.data];
      
        case "ADD_DATA":
        return state;
  
      default:
        return state;
    }
  };
  export default todoDataReducer;
  