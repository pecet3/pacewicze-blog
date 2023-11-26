import React from 'react';

type BasicTileProps = {
    title: string;
};

const BasicTile: React.FC<BasicTileProps> = ({ title }) => {
    return (
        <h2 className="text-3xl">
            {title}
        </h2>
    );
};

export default BasicTile;
