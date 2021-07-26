import React from "react";
import { serviceConfig } from "../applicationSettings";
import { Button } from "react-bootstrap";

class AddCloseFollower extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      listitems: [],
      closeFriends: [],
      removedCloseFriends: [],
    };
    this.handleClick = this.handleClick.bind(this);
    this.handleClickRemove = this.handleClickRemove.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleClick = (e) => {
    const exists = this.state.closeFriends.find(
      (p) => p === e.target.innerText
    );
    if (exists) return;
    var joined = this.state.closeFriends.concat(e.target.innerText);
    this.setState({ closeFriends: joined });
  };

  handleClickRemove = (e) => {
    var array = [...this.state.closeFriends]; // make a separate copy of the array
    var index = array.indexOf(e.target.innerText);
    if (index !== -1) {
      array.splice(index, 1);
      this.setState({ closeFriends: array });
    }
    var joined = this.state.removedCloseFriends.concat(e.target.innerText);
    this.setState({ removedCloseFriends: joined });
  };

  handleSubmit(e) {
    e.preventDefault();
    this.setFollowers();
  }
  renderCloseFollowers() {
    var token = localStorage.getItem("token");
    const requestOpt = {
      method: "GET",
      headers: {
        "Content-Type": "aplication/json",
        Authorization: `Bearer ${token}`,
      },

      credentials: "same-origin",
    };
    fetch(`${serviceConfig.baseURL}/getclosefollowers`, requestOpt)
      .then((response) => response.json())
      .then((responseJson) => {
        var followers = responseJson;
        this.setState({
          closeFriends: followers,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  renderFollowers() {
    var token = localStorage.getItem("token");
    const requestOpt = {
      method: "GET",
      headers: {
        "Content-Type": "aplication/json",
        Authorization: `Bearer ${token}`,
      },

      credentials: "same-origin",
    };
    fetch(`${serviceConfig.baseURL}/getuserwhofollow`, requestOpt)
      .then((response) => response.json())
      .then((responseJson) => {
        var followers = responseJson;
        this.setState({
          listitems: followers,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }
  componentDidMount() {
    this.renderFollowers();
    this.renderCloseFollowers();
  }

  setFollowers() {
    const { closeFriends, removedCloseFriends } = this.state;
    const UsernameDTO = {
      Usernames: closeFriends,
      RemovedUsernames: removedCloseFriends,
    };
    var token = localStorage.getItem("token");
    const requestOpt = {
      method: "POST",
      headers: {
        "Content-Type": "aplication/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(UsernameDTO),
      credentials: "same-origin", // ,'access-control-allow-origin' : '*'
    };

    fetch(`${serviceConfig.baseURL}/setclosefollowers`, requestOpt).then(
      (response) => {
        if (!response.ok) {
          console.log("neuspelo");
          return Promise.reject(response);
        }
        console.log("USPELO");
        return response.json();
      }
    );
  }

  render() {
    return (
      <div className="close-follower-div">
        <React.Fragment>
          <ul className="list-group">
            {this.state.listitems.map((listitem) => (
              <li
                key={listitem}
                className="list-group-item list-group-item-warning list-group-action-variant-warning"
                value={{ listitem }}
                onClick={this.handleClick}
              >
                {listitem}
              </li>
            ))}
          </ul>
        </React.Fragment>
        <div>
          <label></label>
        </div>
        <React.Fragment>
          <ul className="list-group">
            {this.state.closeFriends.map((listitem) => (
              <li
                key={listitem}
                className="list-group-item list-group-item-success "
                value={listitem}
                onClick={this.handleClickRemove}
              >
                {console.log(listitem)}
                {listitem}
              </li>
            ))}
          </ul>
        </React.Fragment>
        <div className="buttonLogin">
          <Button variant="primary" type="submit" onClick={this.handleSubmit}>
            OK
          </Button>
        </div>
      </div>
    );
  }
}
export default AddCloseFollower;
