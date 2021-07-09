import React from 'react';
import './index.css';

export default function Card(props) {
    function movementOver() {
        const card = document.querySelector('.card');
        card.addEventListener("mousemove", (e) => {
            let xAxis = (window.innerWidth / 2 - e.pageX) / 10;
            let yAxis = (window.innerHeight / 2 - e.pageY) / 10;
            card.style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`
        })
    }

    function movementIn() {
        const card = document.querySelector('.card');
        const crypto = document.querySelector('.crypto img');
        const circle = document.querySelector('.circle');
        const title = document.querySelector('.info h1');
        const description = document.querySelector('.info h3');
        card.addEventListener("mouseenter", (e) => {
            card.style.height = '65vh';
            card.style.width = '25rem';
            circle.style.height = '10rem';
            circle.style.width = '10rem';
            crypto.style.width = '8rem';

            card.style.transition = 'all 0.1s ease';
            title.style.transform = 'translateZ(80px)'
            crypto.style.transform = 'translateZ(80px) rotateY(-360deg)'
            description.style.transform = 'translateZ(80px)'
        })
    }

    function movementOut() {
        const card = document.querySelector('.card');
        const crypto = document.querySelector('.crypto img');
        const circle = document.querySelector('.circle');
        const title = document.querySelector('.info h1');
        const description = document.querySelector('.info h3');
        card.addEventListener("mouseleave", (e) => {
            card.style.height = '55vh';
            card.style.width = '20rem';
            circle.style.height = '8rem';
            circle.style.width = '8rem';
            crypto.style.width = '7rem';

            card.style.transition = 'all 0.5s ease';
            card.style.transform = `rotateY(0deg) rotateX(0deg)`;
            circle.style.transition = 'all 0.2s ease';
        
            title.style.transition = 'all 0.5s ease';
            title.style.transform = 'translateZ(0px)'
            crypto.style.transition = 'all 0.5s ease';
            crypto.style.transform = 'translateZ(0px) rotateZ(0deg)'
            description.style.transition = 'all 0.5s ease';
            description.style.transform = 'translateZ(0px)'

        })
    }

    return (
        <div className="card" onMouseOver={movementOver} onMouseLeave={movementOut} onMouseEnter={movementIn}>
            <div className="crypto">
                <div className="circle"></div>
                <img src={props.image} alt={props.name}></img>
            </div>
            <div className="info">
                <h1>{props.name}</h1>
                <h3 style={props.upvotes>=props.downvotes?{color: 'lightgreen'}:{color: 'orangered'}}>Upvotes</h3>
            </div>
        </div>
    );
}