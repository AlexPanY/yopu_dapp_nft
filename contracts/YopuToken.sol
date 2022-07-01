// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./lib/openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./lib/openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "./lib/openzeppelin/contracts/security/Pausable.sol";
import "./lib/openzeppelin/contracts/access/Ownable.sol";
import "./lib/openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "./lib/openzeppelin/contracts/utils/Counters.sol";

/// @custom:security-contact psj474@gmail.com
contract YopuToken is ERC721, ERC721URIStorage, Pausable, Ownable, ERC721Burnable {
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    constructor() ERC721("YopuToken", "YPT") {}

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }

    function _beforeTokenTransfer(address from, address to, uint256 tokenId)
        internal
        whenNotPaused
        override
    {
        super._beforeTokenTransfer(from, to, tokenId);
    }

    // The following functions are overrides required by Solidity.

    function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}