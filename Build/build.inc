#!/bin/bash

set -euxo pipefail

RUN () {
    CMD=$1
    OS="`uname`"
    case $OS in
        Linux*)   
            OSLD_LIBRARY_PATH="LD_LIBRARY_PATH=$2"
            for i in "${@:3}"
            do
                OSLD_LIBRARY_PATH=$OSLD_LIBRARY_PATH:$i
            done
            eval $OSLD_LIBRARY_PATH $CMD
        ;;
        Darwin*)  
            OSLD_LIBRARY_PATH="DYLD_LIBRARY_PATH=$2"
            for i in "${@:3}"
            do
                OSLD_LIBRARY_PATH=$OSLD_LIBRARY_PATH:$i
            done
            eval $OSLD_LIBRARY_PATH $CMD
        ;;
        Windows*) OSLD_LIBRARY_PATH='PATH=%PATH%;' ;;
        *) ;;
    esac
}

# Common initialization
OS="`uname`"
case $OS in
    Linux*) 
        OSEXT='.linux'
        OSLIBEXT='.so'
        OSEXEEXT=
    ;;
    Windows*)
        OSEXT='.exe' 
        OSLIBEXT='.dll'
        OSEXEEXT='.exe'
    ;;
    Darwin*)
        OSEXT='.mac'
        OSLIBEXT='.dylib'
        OSEXEEXT=
    ;;
    *) ;;
esac

export ACT=act$OSEXT
