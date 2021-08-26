import axios from "axios";
import { useState } from "react";
import { useHistory } from "react-router-dom";
import Profile from "./user/Profile";

export default function SearchBar(){

    let history = useHistory();
    const [search,setSearch] = useState('')

    const Search = async () => {
        if (search.includes('#') && search.includes('@')) {
            alert('Please use only one type of search.')
        }
        else if (search.startsWith('@')) {
            history.push('/search/locations/'+search.substring(1,search.length))
        }
        else if (search.startsWith('#')) {
            history.push('/search/tags/'+search.substring(1,search.length))
        }
        else {
            const res = await axios({method:'get',url:'http://localhost:8081/getuserbyusername/'+search});
            if (res.data.username != null) {
                history.push('/profile/'+search);
                window.location.reload();
            }
            else if (window.localStorage.getItem('token') == null) {
                alert('You are not logged in.You will be redirected to login shortly.');
                history.push('/login');
            }
            else {
                alert('There are no users with this username.Usernames are case sensitive.')
            }
        }
    }

    return(
        <div>
        <input type="text" placeholder="@=location,#=tag" value={search} onChange={(e) => setSearch(e.target.value)}/>
        <button onClick={Search}>Search</button>
        </div>
    )
}