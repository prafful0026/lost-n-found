
function Toggler({status, setStatus}) {
    return (
        <div className="flex  border-2 border-primaryBlue rounded-md">
            <button onClick={() => setStatus('lost')} className={`${status==='lost' ? 'bg-primaryBlue text-white' : ' text-primaryBlue'} rounded-md px-4 py-2 transition-all`}>Lost</button>
            <button onClick={() => setStatus('found')} className={`${status==='found' ? 'bg-primaryBlue text-white' : ' text-primaryBlue'} rounded-md px-4 py-2 transition-all`}>Found</button>
        </div>
    )
}

export default Toggler