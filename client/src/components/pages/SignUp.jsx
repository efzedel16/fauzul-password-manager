import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { signUp } from "../../store/action/userAction";

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
        password: password,
      };
      console.log(data);
      dispatch(signUp(data));
    };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <img
            className="mx-auto h-12 w-auto"
            src="https://image.flaticon.com/icons/png/512/4773/4773761.png"
            alt="Workflow"
          />

          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Sign up to your account
          </h2>

          <p className="mt-2 text-center text-sm text-gray-600">
            Or{" "}
            <a
              href="#"
              className="font-medium text-indigo-600 hover:text-indigo-500"
            >
              already have an account? sign in
            </a>
          </p>
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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Password"
                onChange={(e) => {
                  e.preventDefault();
                  setPassword(e.target.value);
                }}
              />
            </div>
          </div>

          {/*<div className="flex justify-center">*/}
          {/*  <div className="text-sm">*/}
          {/*    <Link*/}
          {/*      to="#"*/}
          {/*      className="font-medium text-indigo-600 hover:text-indigo-500"*/}
          {/*    >*/}
          {/*      Forgot your password ?*/}
          {/*    </Link>*/}
          {/*  </div>*/}
          {/*</div>*/}

          <div>
            <button
              type="submit"
              className=" w-full  py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Sign up
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default SignUp;