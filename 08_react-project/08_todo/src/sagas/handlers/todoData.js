import { call, put } from "@redux-saga/core/effects";
import { requestGetData } from "../requests/todoData";
import { setData } from "../../actions";
export function* handleGetData(action) {
  try {
    const response = yield call(requestGetData);
    const data = response;
    // console.log("FROMSAGA",data.data);
    yield put(setData(data.data));
  } catch (error) {
    console.log(error);
  }
}
