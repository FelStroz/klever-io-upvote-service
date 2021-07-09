import React, { useEffect, useState, useRef } from 'react';
import Bitcoin from '../../assets/bitcoin.png'
import Binancecoin from '../../assets/Binancecoin.png'
import Cardano from '../../assets/cardano.png'
import Ethereum from '../../assets/etherium.png'
import Litcoin from '../../assets/Litcoin.png'
import Theter from '../../assets/Theter.png'
import Coin from '../../assets/coin.png'
import Plus from '../../assets/plus.png'
import './home.css';
import Card from '../../components/Card';
import { Link } from "react-router-dom";

function Home() {
  let ref = useRef();

  const [list, setList] = useState([]);
  const [isNewItems, setIsNewItems] = useState(false);
  const [loading, setLoading] = useState(true);

  async function getList() {
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
      setList(res.data);
      setLoading(false);
    }).catch(e => {
      console.log(e)
    })
  }, [])

  function sendUpvote(id) {
    fetch(`http://localhost:8000/crypto/${id}?downvote=0&upvote=1`, {
      "method": "PUT",
      "headers": {}
    })
      .then(response => {
        setIsNewItems(previousState => !previousState);
      })
      .catch(err => {
        console.error(err);
      });
  }

  function sendDownvote(id) {
    fetch(`http://localhost:8000/crypto/${id}?downvote=1&upvote=0`, {
      "method": "PUT",
      "headers": {}
    })
      .then(response => {
        setIsNewItems(previousState => !previousState);
      })
      .catch(err => {
        console.error(err);
      });
  }

  function deleteCrypto(id) {
    fetch(`http://localhost:8000/crypto/${id}`, {
      "method": "DELETE",
      "headers": {}
    })
      .then(response => {
        window.location.replace('http://localhost:3000')
      })
      .catch(err => {
        console.error(err);
      });
  }


  useEffect(() => {
    getList().then((res) => {
      setList(res.data);
    }).catch(e => {
      console.log(e)
    })
  }, [isNewItems])

  return (
    <div>
      <Link to={'/insert'}>
        <div className="insertButton"><text className="buttonText">Insert New</text><img className="addImage" draggable="false" src={Plus} alt="plus"></img></div>
      </Link>
      <div style={{ display: 'flex', flexDirection: 'row', width: '100vw', padding: '2rem' }}>
        {
          !loading ?
            list.map(cryptos => {
              return <Card ref={ref} key={cryptos.id}
                image={
                  (cryptos.name.toLowerCase() === "binance coin") ? Binancecoin
                    : (cryptos.name.toLowerCase() === "cardano") ? Cardano
                      : (cryptos.name.toLowerCase() === "ethereum") ? Ethereum
                        : (cryptos.name.toLowerCase() === "litcoin") ? Litcoin
                          : (cryptos.name.toLowerCase() === "theter") ? Theter
                            : (cryptos.name.toLowerCase() === "bitcoin") ? Bitcoin
                              : Coin}
                name={cryptos.name} upvotes={cryptos.upvote} downvotes={cryptos.downvote} id={cryptos.id} sendUpvote={sendUpvote} sendDownvote={sendDownvote} deleteCrypto={deleteCrypto} />
            })
            :
            <div />
        }
      </div>
    </div>
  );
}

export default Home;