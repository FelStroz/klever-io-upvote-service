import React from 'react';
import './index.css';
import Deletebutton from '../../assets/deletebutton.png';

export default function Card(props) {
    function movementOver() {
        const card = document.getElementsByClassName('card');
        if (card[0])
            card[0].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[0].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
        if (card[1])
            card[1].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[1].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
        if (card[2])
            card[2].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[2].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
        if (card[3])
            card[3].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[3].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
        if (card[4])
            card[4].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[4].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
        if (card[5])
            card[5].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[5].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
        if (card[6])
            card[6].addEventListener("mousemove", (e) => {
                let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
                let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
                card[6].style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
            })
    }

    function movementIn() {
        const card = document.getElementsByClassName('card');
        const crypto = document.getElementsByClassName('image');
        const circle = document.getElementsByClassName('circle');
        const title = document.getElementsByClassName('infoh1');
        const description = document.getElementsByClassName('infoh3');
        const infos = document.getElementsByClassName('infotext');

        if (card[0])
            card[0].addEventListener("mouseenter", (e) => {
                card[0].style.height = '70vh';
                card[0].style.width = '25rem';
                circle[0].style.height = '10rem';
                circle[0].style.width = '10rem';
                crypto[0].style.width = '8rem';

                card[0].style.transition = 'all 0.1s ease';
                title[0].style.transform = 'translateZ(80px)'
                crypto[0].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[0].style.transform = 'translateZ(80px)'
                description[1].style.transform = 'translateZ(80px)'
                infos[0].style.transform = 'translateZ(80px)'
                infos[1].style.transform = 'translateZ(80px)'
            })
        if (card[1])
            card[1].addEventListener("mouseenter", (e) => {
                card[1].style.height = '70vh';
                card[1].style.width = '25rem';
                circle[1].style.height = '10rem';
                circle[1].style.width = '10rem';
                crypto[1].style.width = '8rem';

                card[1].style.transition = 'all 0.1s ease';
                title[1].style.transform = 'translateZ(80px)'
                crypto[1].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[2].style.transform = 'translateZ(80px)'
                description[3].style.transform = 'translateZ(80px)'
                infos[2].style.transform = 'translateZ(80px)'
                infos[3].style.transform = 'translateZ(80px)'
            })
        if (card[2])
            card[2].addEventListener("mouseenter", (e) => {
                card[2].style.height = '70vh';
                card[2].style.width = '25rem';
                circle[2].style.height = '10rem';
                circle[2].style.width = '10rem';
                crypto[2].style.width = '8rem';

                card[2].style.transition = 'all 0.1s ease';
                title[2].style.transform = 'translateZ(80px)'
                crypto[2].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[4].style.transform = 'translateZ(80px)'
                description[5].style.transform = 'translateZ(80px)'
                infos[4].style.transform = 'translateZ(80px)'
                infos[5].style.transform = 'translateZ(80px)'
            })
        if (card[3])
            card[3].addEventListener("mouseenter", (e) => {
                card[3].style.height = '70vh';
                card[3].style.width = '25rem';
                circle[3].style.height = '10rem';
                circle[3].style.width = '10rem';
                crypto[3].style.width = '8rem';

                card[3].style.transition = 'all 0.1s ease';
                title[3].style.transform = 'translateZ(80px)'
                crypto[3].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[6].style.transform = 'translateZ(80px)'
                description[7].style.transform = 'translateZ(80px)'
                infos[6].style.transform = 'translateZ(80px)'
                infos[7].style.transform = 'translateZ(80px)'
            })
        if (card[4])
            card[4].addEventListener("mouseenter", (e) => {
                card[4].style.height = '70vh';
                card[4].style.width = '25rem';
                circle[4].style.height = '10rem';
                circle[4].style.width = '10rem';
                crypto[4].style.width = '8rem';

                card[4].style.transition = 'all 0.1s ease';
                title[4].style.transform = 'translateZ(80px)'
                crypto[4].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[8].style.transform = 'translateZ(80px)'
                description[9].style.transform = 'translateZ(80px)'
                infos[8].style.transform = 'translateZ(80px)'
                infos[9].style.transform = 'translateZ(80px)'
            })
        if (card[5])
            card[5].addEventListener("mouseenter", (e) => {
                card[5].style.height = '70vh';
                card[5].style.width = '25rem';
                circle[5].style.height = '10rem';
                circle[5].style.width = '10rem';
                crypto[5].style.width = '8rem';

                card[5].style.transition = 'all 0.1s ease';
                title[5].style.transform = 'translateZ(80px)'
                crypto[5].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[10].style.transform = 'translateZ(80px)'
                description[11].style.transform = 'translateZ(80px)'
                infos[10].style.transform = 'translateZ(80px)'
                infos[11].style.transform = 'translateZ(80px)'
            })
        if (card[6])
            card[6].addEventListener("mouseenter", (e) => {
                card[6].style.height = '70vh';
                card[6].style.width = '25rem';
                circle[6].style.height = '10rem';
                circle[6].style.width = '10rem';
                crypto[6].style.width = '8rem';

                card[6].style.transition = 'all 0.1s ease';
                title[6].style.transform = 'translateZ(80px)'
                crypto[6].style.transform = 'translateZ(80px) rotateY(-360deg)'
                description[12].style.transform = 'translateZ(80px)'
                description[13].style.transform = 'translateZ(80px)'
                infos[12].style.transform = 'translateZ(80px)'
                infos[13].style.transform = 'translateZ(80px)'
            })

    }

    function movementOut() {
        const card = document.getElementsByClassName('card');
        const crypto = document.getElementsByClassName('image');
        const circle = document.getElementsByClassName('circle');
        const title = document.getElementsByClassName('infoh1');
        const description = document.getElementsByClassName('infoh3');
        const infos = document.getElementsByClassName('infotext');

        if (card[0])
            card[0].addEventListener("mouseleave", (e) => {
                card[0].style.height = '60vh';
                card[0].style.width = '20rem';
                circle[0].style.height = '8rem';
                circle[0].style.width = '8rem';
                crypto[0].style.width = '7rem';

                card[0].style.transition = 'all 0.5s ease';
                card[0].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[0].style.transition = 'all 0.2s ease';

                title[0].style.transition = 'all 0.5s ease';
                title[0].style.transform = 'translateZ(0px)'
                crypto[0].style.transition = 'all 0.5s ease';
                crypto[0].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[0].style.transition = 'all 0.5s ease';
                description[0].style.transform = 'translateZ(0px)'
                description[1].style.transition = 'all 0.5s ease';
                description[1].style.transform = 'translateZ(0px)'
                infos[0].style.transition = 'all 0.5s ease';
                infos[0].style.transform = 'translateZ(0px)'
                infos[1].style.transition = 'all 0.5s ease';
                infos[1].style.transform = 'translateZ(0px)'
            })
        if (card[1])
            card[1].addEventListener("mouseleave", (e) => {
                card[1].style.height = '60vh';
                card[1].style.width = '20rem';
                circle[1].style.height = '8rem';
                circle[1].style.width = '8rem';
                crypto[1].style.width = '7rem';

                card[1].style.transition = 'all 0.5s ease';
                card[1].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[1].style.transition = 'all 0.2s ease';

                title[1].style.transition = 'all 0.5s ease';
                title[1].style.transform = 'translateZ(0px)'
                crypto[1].style.transition = 'all 0.5s ease';
                crypto[1].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[2].style.transition = 'all 0.5s ease';
                description[2].style.transform = 'translateZ(0px)'
                description[3].style.transition = 'all 0.5s ease';
                description[3].style.transform = 'translateZ(0px)'
                infos[2].style.transition = 'all 0.5s ease';
                infos[2].style.transform = 'translateZ(0px)'
                infos[3].style.transition = 'all 0.5s ease';
                infos[3].style.transform = 'translateZ(0px)'
            })
        if (card[2])
            card[2].addEventListener("mouseleave", (e) => {
                card[2].style.height = '60vh';
                card[2].style.width = '20rem';
                circle[2].style.height = '8rem';
                circle[2].style.width = '8rem';
                crypto[2].style.width = '7rem';

                card[2].style.transition = 'all 0.5s ease';
                card[2].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[2].style.transition = 'all 0.2s ease';

                title[2].style.transition = 'all 0.5s ease';
                title[2].style.transform = 'translateZ(0px)'
                crypto[2].style.transition = 'all 0.5s ease';
                crypto[2].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[4].style.transition = 'all 0.5s ease';
                description[4].style.transform = 'translateZ(0px)'
                description[5].style.transition = 'all 0.5s ease';
                description[5].style.transform = 'translateZ(0px)'
                infos[4].style.transition = 'all 0.5s ease';
                infos[4].style.transform = 'translateZ(0px)'
                infos[5].style.transition = 'all 0.5s ease';
                infos[5].style.transform = 'translateZ(0px)'

            })
        if (card[3])
            card[3].addEventListener("mouseleave", (e) => {
                card[3].style.height = '60vh';
                card[3].style.width = '20rem';
                circle[3].style.height = '8rem';
                circle[3].style.width = '8rem';
                crypto[3].style.width = '7rem';

                card[3].style.transition = 'all 0.5s ease';
                card[3].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[3].style.transition = 'all 0.2s ease';

                title[3].style.transition = 'all 0.5s ease';
                title[3].style.transform = 'translateZ(0px)'
                crypto[3].style.transition = 'all 0.5s ease';
                crypto[3].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[6].style.transition = 'all 0.5s ease';
                description[6].style.transform = 'translateZ(0px)'
                description[7].style.transition = 'all 0.5s ease';
                description[7].style.transform = 'translateZ(0px)'
                infos[6].style.transition = 'all 0.5s ease';
                infos[6].style.transform = 'translateZ(0px)'
                infos[7].style.transition = 'all 0.5s ease';
                infos[7].style.transform = 'translateZ(0px)'

            })
        if (card[4])
            card[4].addEventListener("mouseleave", (e) => {
                card[4].style.height = '60vh';
                card[4].style.width = '20rem';
                circle[4].style.height = '8rem';
                circle[4].style.width = '8rem';
                crypto[4].style.width = '7rem';

                card[4].style.transition = 'all 0.5s ease';
                card[4].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[4].style.transition = 'all 0.2s ease';

                title[4].style.transition = 'all 0.5s ease';
                title[4].style.transform = 'translateZ(0px)'
                crypto[4].style.transition = 'all 0.5s ease';
                crypto[4].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[8].style.transition = 'all 0.5s ease';
                description[8].style.transform = 'translateZ(0px)'
                description[9].style.transition = 'all 0.5s ease';
                description[9].style.transform = 'translateZ(0px)'
                infos[8].style.transition = 'all 0.5s ease';
                infos[8].style.transform = 'translateZ(0px)'
                infos[9].style.transition = 'all 0.5s ease';
                infos[9].style.transform = 'translateZ(0px)'

            })
        if (card[5])
            card[5].addEventListener("mouseleave", (e) => {
                card[5].style.height = '60vh';
                card[5].style.width = '20rem';
                circle[5].style.height = '8rem';
                circle[5].style.width = '8rem';
                crypto[5].style.width = '7rem';

                card[5].style.transition = 'all 0.5s ease';
                card[5].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[5].style.transition = 'all 0.2s ease';

                title[5].style.transition = 'all 0.5s ease';
                title[5].style.transform = 'translateZ(0px)'
                crypto[5].style.transition = 'all 0.5s ease';
                crypto[5].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[10].style.transition = 'all 0.5s ease';
                description[10].style.transform = 'translateZ(0px)'
                description[11].style.transition = 'all 0.5s ease';
                description[11].style.transform = 'translateZ(0px)'
                infos[10].style.transition = 'all 0.5s ease';
                infos[10].style.transform = 'translateZ(0px)'
                infos[11].style.transition = 'all 0.5s ease';
                infos[11].style.transform = 'translateZ(0px)'

            })
        if (card[6])
            card[6].addEventListener("mouseleave", (e) => {
                card[6].style.height = '60vh';
                card[6].style.width = '20rem';
                circle[6].style.height = '8rem';
                circle[6].style.width = '8rem';
                crypto[6].style.width = '7rem';

                card[6].style.transition = 'all 0.5s ease';
                card[6].style.transform = `rotateY(0deg) rotateX(0deg)`;
                circle[6].style.transition = 'all 0.2s ease';

                title[6].style.transition = 'all 0.5s ease';
                title[6].style.transform = 'translateZ(0px)'
                crypto[6].style.transition = 'all 0.5s ease';
                crypto[6].style.transform = 'translateZ(0px) rotateZ(0deg)'
                description[12].style.transition = 'all 0.5s ease';
                description[12].style.transform = 'translateZ(0px)'
                description[13].style.transition = 'all 0.5s ease';
                description[13].style.transform = 'translateZ(0px)'
                infos[12].style.transition = 'all 0.5s ease';
                infos[12].style.transform = 'translateZ(0px)'
                infos[13].style.transition = 'all 0.5s ease';
                infos[13].style.transform = 'translateZ(0px)'

            })
    }

    return (
        <div className="card" draggable="false" onMouseOver={movementOver} onMouseLeave={movementOut} onMouseEnter={movementIn}>
            <div className="crypto">
                <div className="circle"></div>
                <img src={props.image} alt={props.name} className="image" draggable="false"></img>
            </div>
            <div className="deleteButton" onClick={() => {props.deleteCrypto(props.id)}}><img className="deleteImage" src={Deletebutton} alt="Delete"></img></div>
            <div className="info">
                <h1 className="infoh1" >{props.name}</h1>
                <h3 className="infoh3" style={{ color: 'lightgreen' }} onClick={() => {props.sendUpvote(props.id)}}>Upvotes</h3>
                <text className="infotext">{props.upvotes}</text>
                <h3 className="infoh3" style={{ color: 'orangered' }} onClick={() => {props.sendDownvote(props.id)}}>Downvotes</h3>
                <text className="infotext">{props.downvotes}</text>

            </div>
        </div>
    );
}