import { createSlice } from "@reduxjs/toolkit";
import { toast } from "react-toastify";

const initialState = {
  user: {
    email: 'testuser@example.com'
  }
}

const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
      loginUser: (state, action) => {
        console.log('login')
      },
      logoutUser: (state) => {
        console.log('logout')
      }
    }
})

export const {loginUser, logoutUser} = userSlice.actions;

export default userSlice.reducer;
