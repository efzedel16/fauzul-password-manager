import API from "../../API/FauzulPasswordManager";

const token = !localStorage.getItem("token")
  ? ""
  : localStorage.getItem("token");

export const createPass = (payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "LOADING" });
      const { data } = await API({
        data: payload,
        header: { Authorization: token },
        method: "POST",
        url: "/password",
      });

      console.log(data);
      return dispatch({ type: "CREATE", payload: data });
    } catch (e) {
      dispatch({ type: "ERROR" });
      console.log(e.response.data);
    }
  };
};

export const updatePass = (id, payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "LOADING" });
      const { data } = await API({
        data: payload,
        header: { Authorization: token },
        method: "PUT",
        url: `/password/${id}`,
      });

      console.log(data);
      return dispatch({ type: "UPDATE", payload: data });
    } catch (e) {
      dispatch({ type: "ERROR" });
      console.log(e.response.data);
    }
  };
};

export const deletePass = (id, payload) => {
  return async (dispatch) => {
    try {
      dispatch({ type: "LOADING" });
      const { data } = await API({
        data: payload,
        header: { Authorization: token },
        method: "DELETE",
        url: `/password/${id}`,
      });

      console.log(data);
      return dispatch({ type: "DELETE", payload: data });
    } catch (e) {
      dispatch({ type: "ERROR" });
      console.log(e.response.data);
    }
  };
};

export const getAllPasswords = () => {
  return async (dispatch) => {
    try {
      dispatch({ type: "LOADING" });
      const { data } = await API({
        header: { Authorization: token },
        method: "GET",
        url: "/passwords",
      });

      console.log(data);
      return dispatch({ type: "GET", payload: data });
    } catch (e) {
      dispatch({ type: "ERROR" });
      console.log(e.response.data);
    }
  };
};