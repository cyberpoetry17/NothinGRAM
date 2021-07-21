import React from 'react';
import { serviceConfig } from "../applicationSettings";

class AddCloseFollower extends React.Component{
    constructor(props) {
        super(props);
    
        this.state = {
          listitems:[],
          closeFriends:[]
        };
        this.handleClick = this.handleClick.bind(this);
        this.handleClickRemove = this.handleClickRemove.bind(this);

        
      }

      handleClick = e =>{
        const exists = this.state.closeFriends.find(p => p === e.target.innerText);
        if (exists) return;
        var joined = this.state.closeFriends.concat(e.target.innerText);
        this.setState({ closeFriends: joined })
        console.log(e.target)  
      }

      handleClickRemove = e =>{
        var array = [...this.state.closeFriends]; // make a separate copy of the array
        var index = array.indexOf(e.target.innerText)
        if (index !== -1) {
          array.splice(index, 1);
          this.setState({closeFriends: array});
        }
      }
      
    
      renderMyData() {
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
                    listitems:  followers
                });
               
              })
              .catch((error) => {
                console.error(error);
              });
          }
          componentDidMount() {
            this.renderMyData();
          }




          render() {
            return (
            <div className="close-follower-div">
               <React.Fragment>
                <ul className="list-group">
                  {this.state.listitems.map(listitem => (
                   
                    <li
                      key={listitem}
                      className="list-group-item list-group-item-warning list-group-action-variant-warning"
                      value={{listitem}}
                      onClick = {this.handleClick}
                      
                    >
                      {console.log(listitem)}
                      {listitem}
                    </li>

                    // <button  key={listitem} type="button" class="list-group-item list-group-item-action" onClick={this.handleClick}>{listitem}</button>
                  ))}
                </ul>
              </React.Fragment>
                <div>
                  <label></label>
                </div>
              <React.Fragment>
              <ul className="list-group">
                {this.state.closeFriends.map(listitem => (
                  <li
                    key={listitem}
                    className="list-group-item list-group-item-success "
                    value={listitem}
                    onClick = {this.handleClickRemove}
                   
                  >
                    {console.log(listitem)}
                    {listitem}
                  </li>
                ))}
              </ul>
              </React.Fragment> 
            </div> 
            );
          }
      } export default AddCloseFollower;
