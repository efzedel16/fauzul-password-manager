const initState = {
  user: null,
  isLoading: false,
  error: null,
};

const userReducer = (state = initState, action) => {
  switch (action.type) {
    case "USER_LOADING":
      return { ...state, isLoading: true };
    case "USER_SIGN_UP":
      return { ...state, isLoading: false, user: action.payload };
    case "USER_SIGN_IN":
      localStorage.setItem("access_token", action.payload.authorization);
      return { ...state, isLoading: false, user: action.payload };
    case "USER_SIGN_OUT":
      localStorage.removeItem("access_token");
      return { ...state, user: null };
    case "USER_ERROR":
      return { ...state, isLoading: false, error: action.payload };
    default:
      return state;
  }
};

export default userReducer;