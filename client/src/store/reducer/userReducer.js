const initState = {
  user: null,
  isLoading: false,
  error: null,
};

const userReducer = (state = initState, action) => {
  switch (action.type) {
    case "LOADING":
      return { ...state, isLoading: true };
    case "SIGN_UP":
      return { ...state, isLoading: false, user: action.payload };
    case "SIGN_IN":
      localStorage.setItem("token", action.payload.authorization);
      return { ...state, isLoading: false, user: action.payload };
    case "SIGN_OUT":
      localStorage.removeItem("token");
      return { ...state, user: null };
    case "ERROR":
      return { ...state, isLoading: false, error: action.payload };
    default:
      return state;
  }
};

export default userReducer;