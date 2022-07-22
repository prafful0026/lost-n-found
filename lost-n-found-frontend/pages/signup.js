import LoginBar from "../components/LoginBar";
import Link from 'next/link'
import {useState} from 'react'
import Logo from "../components/Logo";

export default function Signup() {

    // {
    //     "firstName":"Prafful",
    //     "lastName":"Agrawal",
    //     "email":"avi@gmail.com",
    //     "password":"abcd1234"
    // }

    const initialState = {
        firstName: "",
        lastName: "",
        email: "",
        password: ""
    }
    
    const [userInfo, setUserInfo] = useState(initialState)
    const handleSubmit = (e) => {
        e.preventDefault()
        //handle input
        console.log("Signup Form submitted!", userInfo)
        setUserInfo(initialState)
    }

    const handleChange = (e) => {
        const {name, value} = e.target
        setUserInfo(preValues => ({...preValues, [name]: value}))
    }

    return (
        <div className="text-center ">
            <LoginBar /> {/* If user is not signed up */}
            <Logo />
            <div className="flex flex-col justify-center items-center my-20 px-8">
                <h1 className="font-extrabold  text-4xl mb-8">Join Now</h1>
                <form className="space-y-3 flex flex-col items-center w-full " onSubmit={handleSubmit}>
                    <input name='firstName' onChange={handleChange} className='p-4 max-w-md w-full border-[1px] rounded-md border-gray-200' type="text" placeholder="First Name" />
                    <input name='lastName' onChange={handleChange} className='p-4 max-w-md w-full border-[1px] rounded-md border-gray-200' type="text" placeholder="Last Name" />
                    <input name='email' onChange={handleChange} className='p-4 max-w-md w-full border-[1px] rounded-md border-gray-200' type="email" placeholder="Email" />
                    <input name='password' autoComplete="off" onChange={handleChange} className='p-4 max-w-md w-full border-[1px] rounded-md border-gray-200' type="password" placeholder="Password" />
                    <input type="submit" onClick={handleSubmit} value='Sign Up' className="p-4 text-white max-w-md w-full  bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 btn" />
                </form>
                <hr className='my-6 w-full max-w-lg' />
                <Link href='/login'>
                    <a className="text-primaryBlue">Already have an account? Log in</a>
                </Link>
            </div>
        </div>
    )
}
