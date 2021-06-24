import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { signUp } from "../../store/action/userAction";
import { Link } from "react-router-dom";

const SignUp = () => {
  const dispatch = useDispatch(),
    [fullName, setfullName] = useState(""),
    [address, setAddress] = useState(""),
    [email, setEmail] = useState(""),
    [password, setPassword] = useState(""),
    submitSignUp = (e) => {
      e.preventDefault();
      const data = {
        full_name: fullName,
        address: address,
        email: email,
        pass: password,
      };
      console.log(data);
      dispatch(signUp(data));
    };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <Link to="/">
            <img
              className="mx-auto h-12 w-auto"
              src="https://image.flaticon.com/icons/png/512/4773/4773761.png"
              alt="Workflow"
            />
          </Link>

          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Sign up to secure your credentials
          </h2>
        </div>

        <form className="mt-8 space-y-6" onSubmit={submitSignUp}>
          <input type="hidden" name="remember" defaultValue="true" />
          <div className="rounded-md shadow-sm -space-y-px">
            <div>
              <label htmlFor="full-name" className="sr-only">
                Full name
              </label>
              <input
                id="full-name"
                name="full-name"
                type="text"
                autoComplete="full-name"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                placeholder="Full name"
                onChange={(e) => {
                  e.preventDefault();
                  setfullName(e.target.value);
                }}
              />
            </div>

            <div>
              <label htmlFor="address" className="sr-only">
                Address
              </label>
              <input
                id="address"
                name="address"
                type="text"
                autoComplete="address"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                placeholder="Address"
                onChange={(e) => {
                  e.preventDefault();
                  setAddress(e.target.value);
                }}
              />
            </div>

            <div>
              <label htmlFor="email" className="sr-only">
                Email
              </label>
              <input
                id="email"
                name="email"
                type="email"
                autoComplete="email"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                placeholder="Email"
                onChange={(e) => {
                  e.preventDefault();
                  setEmail(e.target.value);
                }}
              />
            </div>

            <div>
              <label htmlFor="password" className="sr-only">
                Password
              </label>
              <input
                id="password"
                name="password"
                type="password"
                autoComplete="password"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                placeholder="Password"
                onChange={(e) => {
                  e.preventDefault();
                  setPassword(e.target.value);
                }}
              />
            </div>
          </div>

          <div>
            <button
              type="submit"
              className=" w-full  py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            >
              Sign up
            </button>
          </div>
        </form>

        <p className="mt-2 text-center text-sm text-gray-600">
          already have an account?{" "}
          <Link
            className="font-medium text-green-600 hover:text-green-500"
            to="/signin"
          >
            sign in
          </Link>
        </p>
      </div>
    </div>
  );
};

export default SignUp;