import Link from "next/link"

function LoginBar() {
  return (
    <div className="top-0 h-12 w-full space-x-4 border-b-[1px] border-black border-opacity-10 flex justify-center items-center">
        <Link href="/login">
          <a className="text-primaryBlue">Login</a>
        </Link>
        <Link href="/signup" passHref>
          <div className="bg-primaryBlue text-white py-1 px-3 rounded-md">
            <a>Sign Up</a> 
          </div>
        </Link>
    </div>
  )
}

export default LoginBar