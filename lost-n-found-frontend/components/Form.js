import { useState } from "react"
import Toggler from "./Toggler"

function Form({ stateObject, statusObject }) {
    const isPostCreation = !!statusObject
    const { state, initialState, setState } = stateObject
    const { status, setStatus } = statusObject
    const [fileName, setFileName] = useState([])
    const handleSubmit = (e) => {
        e.preventDefault()
        //access state
        console.log("Post submitted!", post)
        setState(initialState)
    }
    const handleChange = (e) => {
        const { name, value } = e.target
        setState(preValue => ({
            ...preValue,
            [name]: value
        }))
    }

    function showFileName() {
        var fileList = document.getElementById('inputFile')
        // console.log(fileList?.files.item(0).name)
        for(let i = 0; i<fileList?.files?.length; i++){
            setFileName(state => [...state, fileList?.files?.item(i).name])
        }
    }

    return (
        <>
            {/* <h1 className="font-bold text-3xl md:text-4xl text-center my-3">{status?'Create Post':'Claim Item'}</h1>
            <div className="w-full flex justify-center">
                <Toggler status={status} setStatus={setStatus} />
            </div>
            <p className='text-center my-3 italic opacity-50'>This information will be shared publicly ⬇️</p> */}
            <div className="w-full flex justify-center py-4">
                <form onSubmit={handleSubmit} className='font-medium flex flex-col items-center w-full max-w-md'>
                    <label className="opacity-80 w-full">Title:</label>
                    <input onChange={handleChange} required={true} name='title' placeholder={status === 'found' ? 'Eg."Found red color wallet"' : 'Eg."Lost a red color wallet"'} type='text' className='p-4 mt-2 mb-3 max-w-md w-full bg-lightGray rounded-md ' />
                    <label className="opacity-80 w-full">Description:</label>
                    <textarea onChange={handleChange} required={true} name="description" rows={4} className="p-4 mt-2 mb-3 max-w-md w-full bg-lightGray rounded-md " />
                    <label className="opacity-80 w-full">Email:</label>
                    <input onChange={handleChange} name='email' type='email' className='p-4 mt-2 mb-3 max-w-md w-full bg-lightGray rounded-md ' />
                    <label className="opacity-80 w-full">Address:</label>
                    <input onChange={handleChange} required={true} name='address' type='text' className='p-4 mt-2 mb-3 max-w-md w-full bg-lightGray rounded-md ' />
                    <label className="opacity-80 w-full">Mobile Number:</label>
                    <input onChange={handleChange} required={true} name='mobileNumber' type='tel' className='p-4 mt-2 mb-3 max-w-md w-full bg-lightGray rounded-md ' />
                    <label className="opacity-80 w-full">Image of the {status === 'found' ? "Item you found" : "Lost Item"}:</label>
                    <input name='photo' id='inputFile' hidden type="file" multiple accept="image/png, image/jpeg" className="" onChange={showFileName} />
                    <div className="w-full flex gap-2 mt-4 m-3">
                        {
                            fileName && (
                                fileName.map(file => <label key={file} htmlFor="inputFile" className="bg-lightGray flex-wrap p-4 text-center  cursor-pointer rounded-md opacity-80">{file}</label>)
                            )
                        }
                        <label htmlFor="inputFile" className="bg-lightGray p-4 text-center w-32 cursor-pointer rounded-md opacity-80">Choose file</label>
                    </div>
                    <input type="submit" required={status === 'found' ? true : false} value='Post' className="bg-primaryBlue cursor-pointer w-full font-medium text-white mt-4 p-4 rounded-md" />
                </form>
            </div>
        </>
    )
}

export default Form
