import React, {useEffect, useState, useRef} from 'react';
import Bitcoin from '../../assets/bitcoin.png'
import './home.css';
import Card from '../../components/Card';

function Home() {
  let ref = useRef();

  const [list, setList] = useState({});
 
  async function getList(){
    return fetch("http://localhost:8000/crypto", {
      "method": "GET",
      "headers": {}
    })
    .then(response => {
      return response.json();
    })
    .catch(err => {
      return (err);
    });
}

  useEffect(() => {
    getList().then((res) => {
      console.log(res);
     // setList(res);
    }).catch(e => {
      console.log(e)
    })
  },[])

  return (
    <div style={{display:'flex', flexDirection: 'row'}}>
      <Card ref={ref} image={Bitcoin} name={"Bitcoin"} upvotes={1} downvotes={1} />
    </div>
  );
}

export default Home;
