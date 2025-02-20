import React from 'react';
import '../main.css';
import './search.css';


const SearchPage: React.FC = () => {
  return (
    <>
      <div id="astheader">
      </div>
      
      <div id="searcharea">
        <input name="search" placeholder="タグを検索" />
        <select name="sex" defaultValue="none">
          <option value="none">言語を選択</option>
          <option value="javascript">javascript</option>
        </select>
        <input type="submit" value="検索" />
      </div>
      
      <div id="main">
      </div>
    </>
  );
};

export default SearchPage;
