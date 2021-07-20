import React, { useState } from "react";
import DatePicker from "react-datepicker";

// import required css from library
import "react-datepicker/dist/react-datepicker.css";

// main implementation. using selected and onChange as main props to get and change the selected date value



const DatePick = () => {
  const [this.props.startDate, setStartDate] = useState(new Date());
  // const min = new Date(1900, 1, 1);
  // const max = new Date(2021, 8, 10);
  // const maxDate = new Date(2021, 12, 12);

  return (
    <DatePicker
      selected={this.props.startDate}
      onChange={(date) => setStartDate(date)}
      isClearable
      placeholderText="I have been cleared!"
    />
  );
  
}

export default DatePick;
