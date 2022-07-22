import Image from 'next/image'
import Link from 'next/link'

function BottomNav() {
    return (
        <div className="fixed left-1/2 transform -translate-x-[50%] py-3 rounded-t-3xl font-medium bottom-0 w-full max-w-lg bg-white z-10 drop-shadow-2xl flex justify-around items-center">
            <Link href='#' >
                <a>Log out</a>
            </Link>
            <Link href='/posts/createPost' passHref>
                <div className='cursor-pointer'>
                    <div className='absolute -top-12 left-1/2 transform -translate-x-[45%]'>
                        <Image src='/Plus.svg' height={70} width={70} alt='New Post' />
                    </div>
                    <a>New Post</a>
                </div>
            </Link>
            <Link href='#'>
                <a>Profile</a>
            </Link>
        </div>
    )
}

export default BottomNav

// type Post struct {
// 	Id          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
// 	Title*      string               `json:"title,omitempty" validate:"required"`
// 	Description* string               `json:"description,omitempty"`
// 	User        primitive.ObjectID   `json:"user,omitempty"`
// 	Email       string               `json:"email,omitempty" validate:"email"`
// 	Address*     string               `json:"address,omitempty" validate:"required"`
// 	PhoneNumber* string               `json:"phoneNumber,omitempty" validate:"required"`
// 	ImageUrls   []string             `json:"imageUrls,omitempty" validate:"required"`
// 	Status*      string               `json:"status,omitempty" validate:"required, eq=LOST|eq=FOUND"`
// 	Claims      []primitive.ObjectID `json:"claims,omitempty"`
// 	CreatedAt   time.Time            `json:"createdAt,omitempty"`
// }