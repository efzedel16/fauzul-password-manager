import thunk from "redux-thunk";
import rootReducer from "./rootReducer";
import { applyMiddleware, createStore } from "redux";

const store = createStore(rootReducer, applyMiddleware(thunk));

export default store;