import { combineReducers } from "redux";
import user from "./reducer/userReducer";
import pass from "./reducer/passReducer";

const rootReducer = combineReducers({
  user,
  pass,
});

export default rootReducer;