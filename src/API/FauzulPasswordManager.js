import Axios from "axios";

const API = Axios.create({
  baseURL: "https://fauzul-password-manager.herokuapp.com",
});

export default API;