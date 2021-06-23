import API from "../../API/FauzulPasswordManager";

export const userSignUp = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "USER_LOADING" });
      const { data } = await API({
        method: "POST",
        url: "/user/signup",
        data: payload,
      });

      console.log(data);
      return dispatch({ type: "USER_SIGN_UP", payload: data });
    } catch (e) {
      console.log(e.response);
    }
  };
};