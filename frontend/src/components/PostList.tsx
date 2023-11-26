import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface Post {
    id: string;
    user_id: string;
    title: string;
    content: string;
    image_url: string;
    created_at: string;
}

const PostList: React.FC = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async (): Promise<void> => {
            try {
                setLoading(true); // Set loading to true before making the request
                const response = await axios.get('http://srv23.mikr.us:30111/api/post');
                setPosts(response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            } finally {
                setLoading(false); // Set loading to false after the request completes or encounters an error
            }
        };

        fetchData().then(() => {
            console.log('Data fetched successfully! You can perform additional actions here.');
        });
    }, []);

    return (
        <div>
            <h1>Post List</h1>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <ul>
                    {posts.map(post => (
                        <li key={post.id}>
                            <h2>{post.title}</h2>
                            <p>{post.content}</p>
                            <p>Created at: {post.created_at}</p>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default PostList;
