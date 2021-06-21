import API from "../../API/FauzulPasswordManager";

export const userSignUp = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "USER_LOADING" });
      const { data } = await API({
        method: "POST",
        url: "/USER/signup",
        data: payload,
      });

      console.log(data);

      return dispatch({ type: "USER_SIGN_UP", payload: data });
    } catch (e) {
      console.log(e.response);
    }
  };
};

export const userSignIn = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "USER_LOADING" });
      const { data } = await API({
        method: "POST",
        url: "/user/signin",
        data: payload,
      });

      localStorage.setItem("access_token", data.authorization);
      return dispatch({ type: "USER_SIGN_IN", payload: data });
    } catch (e) {
      console.log(e.response);
    }
  };
};

export const userSignOut = () => ({
  type: "USER_SIGN_OUT",
});