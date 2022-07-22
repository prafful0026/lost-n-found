import { useState } from "react"
import Form from "../../components/Form"
import Logo from "../../components/Logo"

function CreatePost() {
    const initialState = {
        title: '',
        description: '',
        email: '',
        address: '',
        mobileNumber: '',
    }
    const [status, setStatus] = useState('found')
    const [post, setPost] = useState(initialState)

    // const handleSubmit = (e) => {
    //     e.preventDefault()
    //     //handleinput
    //     console.log("Post submitted!", post)
    //     setPost(initialState)
    // }
    // const handleChange = (e) => {
    //     const { name, value } = e.target
    //     setPost(preValue => ({
    //         ...preValue,
    //         [name]: value
    //     }))
    // }
    return (
        <div className="px-4">
            {/* Header */}
            <div className="absolute left-0">
                <Logo />
            </div>
            <Form statusObject={{status, setStatus}} stateObject={{initialState, post, setPost}} />
        </div>
    )
}

export default CreatePost