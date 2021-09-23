import axios from 'axios'
export function requestGetData() {
    return axios.request({
        method: "get",
        url:"http://6146ecde65467e00173849b9.mockapi.io/todoApi/task"
    })
} 