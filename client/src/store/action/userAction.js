import API from "../../API/FauzulPasswordManager";

export const signUp = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "LOADING" });
      const { data } = await API({
        data: payload,
        method: "POST",
        url: "/signup",
      });

      console.log(data);
      return dispatch({ payload: data, type: "SIGN_UP" });
    } catch (e) {
      console.log(e.response);
    }
  };
};

export const signIn = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "LOADING" });
      const { data } = await API({
        data: payload,
        method: "POST",
        url: "/signin",
      });

      console.log(data);
      return dispatch({ payload: data, type: "SIGN_IN" });
    } catch (e) {
      console.log(e.response);
    }
  };
};