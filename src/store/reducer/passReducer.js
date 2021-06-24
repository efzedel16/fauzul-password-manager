const initState = {
  passwords: [],
  password: null,
  isLoading: false,
  error: null,
};

const passReducer = (state = initState, action) => {
  switch (action.type) {
    case "LOADING":
      return { ...state, isLoading: true };
    case "GET":
      return { ...state, isLoading: false, passwords: action.payload };
    case "CREATE":
      return { ...state, isLoading: false, password: action.payload };
    case "UPDATE":
      return { ...state, isLoading: false, password: action.payload };
    case "DELETE":
      return { ...state, isLoading: false, password: null };
    case "ERROR":
      return { ...state, isLoading: false };
    default:
      return state;
  }
};

export default passReducer;