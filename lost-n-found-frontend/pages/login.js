import LoginBar from "../components/LoginBar"
import Link from 'next/link'
import {useState} from 'react'

function Login() {
    const initialState = {
        email: "",
        password: "",
    }
    const [userInfo, setUserInfo] = useState(initialState)
    
    const handleSubmit = (e) => {
        e.preventDefault()
        //handleinput
        console.log("Login credentials submitted!", userInfo)
        setUserInfo(initialState)
    }
    const handleChange = (e) => {
        const {name, value} = e.target
        setUserInfo(preValue => ({
            ...preValue,
            [name]: value
        }))
    } 

    return (
        <div className="text-center ">
            <LoginBar /> {/* If user is not signed up */}
            <div className="flex flex-col justify-center items-center my-20 px-8">
                <h1 className="font-extrabold  text-4xl mb-8">Login</h1>
                <form onSubmit={handleSubmit} className="space-y-3 flex flex-col items-center w-full ">
                    <input onChange={handleChange} name='email' className='p-4 max-w-md w-full border-[1px] rounded-md border-gray-200' type="email" placeholder="Email" />
                    <input onChange={handleChange} name='password' autoComplete="off" className='p-4 max-w-md w-full border-[1px] rounded-md border-gray-200' type="password" placeholder="Password" />
                    <input onClick={handleSubmit}  type="submit" value='Login' className="p-4 text-white max-w-md w-full  bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 btn" />
                </form>
                <hr className='my-6 w-full max-w-lg' />
                <Link href='/signup'>
                    <a className="text-primaryBlue">Don{"'"}t have an account? Sign Up</a>
                </Link>
            </div>
        </div>
    )
}

export default Login