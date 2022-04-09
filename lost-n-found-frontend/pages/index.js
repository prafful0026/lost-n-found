import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'
import LoginBar from '../components/LoginBar'
import Logo from '../components/Logo'
import Footer from '../components/Footer'
import { useState } from 'react/cjs/react.production.min'

export default function Home() {
  const [name, setName] = useState
  return (
    <div>
      <Head>
        <title>Lost & Found App</title>
        <meta name="description" content="Lost & Found App" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className='container text-center'>
        <LoginBar /> {/* If user is not signed up */}
        <Logo />
        <div className=' flex flex-col justify-center items-center space-y-4 my-24'>
          <h1 className='text-4xl font-extrabold bg-gradient-to-r text-transparent bg-clip-text from-indigo-500 via-purple-500 to-pink-500'>Find the lost,<br />Help the losers</h1>
          <h2 className='font-bold text-2xl'>Lost & Found App</h2>

          <div className='flex space-x-3 px-8'>
            <Link href='/posts' passHref>
              <div className='btn bg-black border-[1px] border-black w-20'>
                <a className='text-white font-medium'>Find</a>
              </div>
            </Link>
            <Link href='/posts' passHref>
              <div className='btn border-[1px] border-black w-20'>
                <a className=' font-medium'>Found</a>
              </div>
            </Link>
          </div>
        </div>
        <div className='px-8 space-y-10 my-24 bg-[#F9F9F9] py-10'>
          <div>
            <Image src='/care.svg' height={30} width={70} objectFit='cover' alt='' />
            <p className='font-medium opacity-60 text-sm'>We care about you and your things. We know how much these things mean for you. Its painfull when you lose any. But don{"'"}t worry we are there to help you out.</p>
          </div>
          <div>
            <Image src='/image.svg' height={50} width={50} objectFit='cover' alt='' />
            <p className='font-medium opacity-60 text-sm'>Found a lost item, but unable to find the owner? We will help you find the right owner of the item. You know, Helping others is good.</p>
          </div>
        </div>
        <Footer />
      </div>
    </div>
  )
}
