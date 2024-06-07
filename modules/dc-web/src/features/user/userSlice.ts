import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export type User = {
  userId: string,
  jwt: string,
}

const initialState = {
  name: "user",

};

const userSlice = createSlice({
  name:"user slice",
  initialState,
  reducers: {

  },
})

export default userSlice.reducer;
