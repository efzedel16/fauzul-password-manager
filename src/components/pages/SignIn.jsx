import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { signIn } from "../../store/action/userAction";
import { Link } from "react-router-dom";

const SignIn = () => {
  const dispatch = useDispatch(),
    [email, setEmail] = useState(""),
    [password, setPassword] = useState(""),
    submitSignUp = (e) => {
      e.preventDefault();
      const data = {
        email: email,
        pass: password,
      };
      console.log(data);
      dispatch(signIn(data));
    };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <Link to="/">
            <img
              className="mx-auto h-12 w-auto"
              src="https://image.flaticon.com/icons/png/512/4773/4773761.png"
              alt="Fauzul Password Manager's Icon"
            />
          </Link>

          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Sign in to your account
          </h2>
        </div>

        <form className="mt-8 space-y-6" onSubmit={submitSignUp}>
          <input type="hidden" name="remember" defaultValue="true" />
          <div className="rounded-md shadow-sm -space-y-px">
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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
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
              Sign in
            </button>
          </div>
        </form>

        <p className="mt-2 text-center text-sm text-gray-600">
          don't have an account yet?{" "}
          <Link
            className="font-medium text-green-600 hover:text-green-500"
            to="/signup"
          >
            sign up
          </Link>
        </p>
      </div>
    </div>
  );
};

export default SignIn;