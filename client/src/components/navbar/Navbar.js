import React, {Component} from 'react';
import { LinkContainer } from "react-router-bootstrap";
import './Navbar.css'

class Navbar extends Component {
    state = { clicked: false }

    // Sets dropdown icon to clicked or not clicked
    handleClick = () => {
        this.setState({ clicked: !this.state.clicked })
    }

    render() {
        return(
            <nav className="NavbarItems">
                <LinkContainer to="/">
                    <h1 className="navbar-logo">Nill Book Club <i className='fas fa-quran'></i></h1>
                </LinkContainer>
                <div className="menu-icon" onClick={this.handleClick}>
                    <i className={this.state.clicked ? 'fas fa-times' : 'fas fa-bars'}></i>
                </div>

                <ul className={this.state.clicked ? 'nav-menu active' : 'nav-menu'}>
                    <li>
                        <LinkContainer to="/">
                            <a className="nav-links" href="/">Home</a>
                        </LinkContainer>
                    </li>
                    <li>
                        <LinkContainer to="/login">
                            <a className="nav-links" href="/login">Login</a>
                        </LinkContainer>
                    </li>
                    
                </ul>
            </nav>
        )
    }
}

export default Navbar
