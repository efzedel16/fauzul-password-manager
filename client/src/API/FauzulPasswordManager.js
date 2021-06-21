import Axios from "axios";

const API = Axios.create({
  baseURL: "http://localhost:8080",
});

export default API;